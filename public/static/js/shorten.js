'use strict';

document.addEventListener('DOMContentLoaded', () => {
    const { copyBtnElem, submitBtnElem, formElem, urlInputElem, errElem } = {
        copyBtnElem: document.getElementById("shortenCopyButton"),
        submitBtnElem: document.getElementById("shortenSubmitButton"),
        formElem: document.getElementById("shortenForm"),
        urlInputElem: document.getElementById("shortenURLInput"),
        errElem: document.getElementById("errorMsg"),
    };

    const err = (msg) => {
        errElem.innerHTML = msg;
        errElem.style.display = "block";
        setTimeout(() => {
            errElem.style.display = "none";
            errElem.innerHTML = "";
        }, 1000);
    }
    const shortenURLInputTrigger = (e) => {
        e.preventDefault();
        copyBtnElem.style.display = "none";
        submitBtnElem.style.display = "block";
        urlInputElem.removeEventListener("change", shortenURLInputTrigger);
        urlInputElem.removeEventListener("input", shortenURLInputTrigger);
    }
    formElem.addEventListener("submit", (e) => {
        e.preventDefault();
        let xhr = new XMLHttpRequest();
        xhr.open("POST", '/api/shorten');
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.onload = () => {
            let resp = JSON.parse(xhr.responseText);
            if (xhr.status != 200) {
                err(xhr.statusText);
                return;
            } else if (resp.errcode) {
                err(resp.errmsg);
                return;
            }
            urlInputElem.value = resp.data
            copyBtnElem.style.display = "block";
            submitBtnElem.style.display = "none";
            urlInputElem.addEventListener("change", shortenURLInputTrigger);
            urlInputElem.addEventListener("input", shortenURLInputTrigger);
        };
        xhr.send("url=" + urlInputElem.value);
    });
    copyBtnElem.addEventListener("click", e => {
        e.preventDefault();
        if (!navigator.clipboard) {
            err("unable to copy");
            return;
        }
        const { state } = navigator.permissions.query({
            name: "clipboard-write",
        });
        if (state == "denied") {
            err("clipboard permission is denied.");
            return;
        }
        navigator.clipboard.writeText(urlInputElem.value).then(() => {
            copyBtnElem.classList.add("copied");
            let buttonCopyTimer;
            clearTimeout(buttonCopyTimer);
            buttonCopyTimer = setTimeout(() => {
                copyBtnElem.classList.remove("copied");
            }, 1000);
        }).catch(() => {
            err("unable to copy");
        });
    })
});