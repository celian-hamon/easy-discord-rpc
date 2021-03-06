package main

var mainPage = `
<!DOCTYPE html>
<html>

<head>
    <meta charset='utf-8'>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'>
    <title>webapp</title>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet"
        integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.8.0/font/bootstrap-icons.css">
    <script>

    </script>
    <style>
        .spawn {
            animation: spawn 0.75s;
        }

        .animated-border-button {
            background-color: #444b52;
            border: none;
            color: #ffffff;
            outline: none;
            padding: 12px 40px 10px;
            position: relative;
        }

        .animated-border-button:before,
        .animated-border-button:after {
            border: 0 solid transparent;
            transition: all 0.4s;
            content: '';
            height: 0;
            position: absolute;
            width: 24px;
        }

        .animated-border-button:before {
            border-top: 2px solid rgb(165, 127, 133);
            right: 0;
            top: -4px;
        }

        .animated-border-button:after {
            border-bottom: 2px solid rgb(24, 103, 182);
            bottom: -4px;
            left: 0;
        }

        .animated-border-button:hover {
            background-color: #212529;

        }

        .animated-border-button:hover:before,
        .animated-border-button:hover:after {
            width: 100%;
        }

        @keyframes spawn {
            0% {
                opacity: 0;
                transform: translateY(50%);
            }

            100% {
                opacity: 1;
                transform: translateY(0);
            }
        }

        body {
            --s: 25vmin;
            --p: calc(var(--s) / 2);
            --c1: rgb(165, 127, 133);
            --c2: rgb(24, 103, 182);
            --c3: #212529;
            --bg: var(--c3);
            --d: 8000ms;
            --e: cubic-bezier(0.76, 0, 0.24, 1);

            background-color: var(--bg);
            background-image:
                linear-gradient(45deg, var(--c1) 25%, transparent 25%),
                linear-gradient(-45deg, var(--c1) 25%, transparent 25%),
                linear-gradient(45deg, transparent 75%, var(--c2) 75%),
                linear-gradient(-45deg, transparent 75%, var(--c2) 75%);
            background-size: var(--s) var(--s);
            background-position:
                calc(var(--p) * 1) calc(var(--p) * 0),
                calc(var(--p) * -1) calc(var(--p) * 1),
                calc(var(--p) * 1) calc(var(--p) * -1),
                calc(var(--p) * -1) calc(var(--p) * 0);
            animation:
                color var(--d) var(--e) infinite,
                position var(--d) var(--e) infinite;
        }

        @keyframes color {

            0%,
            25% {
                --bg: var(--c3);
            }

            26%,
            50% {
                --bg: var(--c1);
            }

            51%,
            75% {
                --bg: var(--c3);
            }

            76%,
            100% {
                --bg: var(--c2);
            }
        }

        @keyframes position {
            0% {
                background-position:
                    calc(var(--p) * 1) calc(var(--p) * 0),
                    calc(var(--p) * -1) calc(var(--p) * 1),
                    calc(var(--p) * 1) calc(var(--p) * -1),
                    calc(var(--p) * -1) calc(var(--p) * 0);
            }

            25% {
                background-position:
                    calc(var(--p) * 1) calc(var(--p) * 4),
                    calc(var(--p) * -1) calc(var(--p) * 5),
                    calc(var(--p) * 1) calc(var(--p) * 3),
                    calc(var(--p) * -1) calc(var(--p) * 4);
            }

            50% {
                background-position:
                    calc(var(--p) * 3) calc(var(--p) * 8),
                    calc(var(--p) * -3) calc(var(--p) * 9),
                    calc(var(--p) * 2) calc(var(--p) * 7),
                    calc(var(--p) * -2) calc(var(--p) * 8);
            }

            75% {
                background-position:
                    calc(var(--p) * 3) calc(var(--p) * 12),
                    calc(var(--p) * -3) calc(var(--p) * 13),
                    calc(var(--p) * 2) calc(var(--p) * 11),
                    calc(var(--p) * -2) calc(var(--p) * 12);
            }

            100% {
                background-position:
                    calc(var(--p) * 5) calc(var(--p) * 16),
                    calc(var(--p) * -5) calc(var(--p) * 17),
                    calc(var(--p) * 5) calc(var(--p) * 15),
                    calc(var(--p) * -5) calc(var(--p) * 16);
            }
        }

        @media (prefers-reduced-motion) {
            body {
                animation: none;
            }
        }

        html {
            overflow: hidden;
        }

        .bg-image-vertical {
            position: relative;
            overflow: hidden;
            background-repeat: no-repeat;
            background-position: right center;
            background-size: auto 100%;
        }


        @media (min-width: 1025px) {
            .h-custom-2 {
                height: 100%;
            }
        }

        #time {
            margin-bottom: 30px;
        }
    </style>
</head>

<body>
    <section class="vh-100 gradient-custom">
        <div class="container py-5 h-100">
            <div class="row d-flex justify-content-center align-items-center h-100">
                <div class="col-12 col-md-8 col-lg-6 col-xl-5">
                    <div class="card bg-dark text-white spawn" style="border-radius: 1rem;">
                        <div class="card-body p-5 text-center">

                            <div class="mb-md-2 mt-md-4 pb-3">
                                <form>

                                    <h1 class="fw-normal mb-3 pb-3" style="letter-spacing: 1px;"><strong>CUSTOM
                                            RPC</strong></h1>
                                    <div class="row">
                                        <div class="col form-outline form-dark mb-4">
                                            <input type="number" id="rpcAppId" name="rpcAppId"
                                                class="form-control form-control-lg " disabled
                                                value="935173881256345751" />
                                            <div class="form-check form-switch">
                                                <label class="form-label" for="rpcAppId">APP id</label>
                                                <input class="form-check-input" type="checkbox"
                                                    id="flexSwitchCheckChecked" onclick="test()" checked>
                                            </div>
                                        </div>

                                        <div class="col form-outline form-dark mb-4">
                                            <input type="text" id="rpcState" name="rpcState"
                                                class="form-control form-control-lg" />

                                            <label class="form-label" for="rpcState">State</label>

                                        </div>

                                    </div>
                                    <div class="row">
                                        <div class="col">
                                            <input type="text" id="rpcDetails" name="rpcDetails"
                                                class="form-control form-control-lg" />
                                            <label for="rpcDetails">Details</label>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col">
                                            <input type="text" id="rpcLargeImg" name="rpcLargeImg"
                                                class="form-control form-control-lg" onchange="checkurl(this)" />
                                            <label for="rpcDetails">Url grande image</label>
                                        </div>
                                        <div class="col">
                                            <input type="text" id="rpcLargeTxt" name="rpcLargeTxt"
                                                class="form-control form-control-lg" />
                                            <label for="rpcDetails">Texte grande image</label>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col">
                                            <input type="text" id="rpcSmallImg" name="rpcSmallImg"
                                                class="form-control form-control-lg" onchange="checkurl(this)" />
                                            <label for="rpcDetails">Url petite image</label>
                                        </div>
                                        <div class="col">
                                            <input type="text" id="rpcSmallTxt" name="rpcSmallTxt"
                                                class="form-control form-control-lg" />
                                            <label for="rpcDetails">Texte petite image</label>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col">
                                            <input type="number" id="rpcPartyNum" name="rpcPartyNum"
                                                class="form-control form-control-lg" />
                                            <label for="rpcDetails">Party Num</label>
                                        </div>
                                        <div class="col">
                                            <input type="number" id="rpcPartyMax" name="rpcPartyMax"
                                                class="form-control form-control-lg" />
                                            <label for="rpcDetails">Party Max</label>
                                        </div>
                                    </div>
                                    <div class="row" id="time">
                                        <div class="col">
                                            <input type="time" id="rpcStart" name="rpcStart"
                                                class="form-control form-control-lg" />
                                            <label for="rpcDetails">Rpc Start</label>
                                        </div>
                                        <div class="col">
                                            <input type="time" id="rpcEnd" name="rpcEnd"
                                                class="form-control form-control-lg" />
                                            <label for="rpcDetails">Rpc End</label>
                                        </div>
                                    </div>

                                    <div class="row">
                                        <div class="pt-1 mb-4">
                                            <button class="btn-lg btn-primary animated-border-button" type="button"
                                                onclick="send()"><strong>SUBMIT</strong></button>
                                        </div>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
    </section>
    <script>
        function send() {
            // Sending and receiving data in JSON format using POST method
            //
            var xhr = new XMLHttpRequest();
            var url = "http://localhost:8080/setRPC";
            xhr.open("POST", url, true);
            xhr.setRequestHeader("Content-Type", "application/json");
            var data = JSON.stringify({
                "rpcAppId": document.getElementById("rpcAppId").value,
                "rpcState": document.getElementById("rpcState").value,
                "rpcDetails": document.getElementById("rpcDetails").value,
                "rpcLargeImg": document.getElementById("rpcLargeImg").value,
                "rpcLargeTxt": document.getElementById("rpcLargeTxt").value,
                "rpcSmallImg": document.getElementById("rpcSmallImg").value,
                "rpcSmallTxt": document.getElementById("rpcSmallTxt").value,
                "rpcPartyNum": document.getElementById("rpcPartyNum").value,
                "rpcPartyMax": document.getElementById("rpcPartyMax").value,
                "rpcStart": document.getElementById("rpcStart").value,
                "rpcEnd": document.getElementById("rpcEnd").value
            });
            console.log(data);
            xhr.send(data);
        }

        function test() {
            if (document.getElementById("flexSwitchCheckChecked").checked) {
                document.getElementById("rpcAppId").disabled = true;
                document.getElementById("rpcAppId").value = "935173881256345751";
            } else {
                document.getElementById("rpcAppId").disabled = false;
                document.getElementById("rpcAppId").value = "";

            }
        }
        function checkurl(element) {
            if (element.value.match(/\.(jpeg|jpg|gif|png)$/)) {
                element.classList.remove("is-invalid");
                element.classList.add("is-valid");
            } else {
                element.classList.remove("is-valid");
                element.classList.add("is-invalid");
            }
        }

    </script>
</body>

</html>
`
