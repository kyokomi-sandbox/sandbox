import http from "k6/http"
import {sleep} from "k6"

import {url} from "./config.js"


export default function () {
  //http.post(url(`/isuconp.v1.IsuconpService/Initialize`), {
  http.get(url(`/initialize`), {
    timeout: "10s",
  })
  sleep(1)
}
