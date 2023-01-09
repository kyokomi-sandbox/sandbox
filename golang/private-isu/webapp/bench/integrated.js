import initialize from "./initialize.js"
import comment from "./comment.js"
import postImage from "./postimage.js"

export {initialize, comment, postImage}

export const options = {
  scenarios: {
    initialize: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 1,
      exec: "initialize",
      maxDuration: "10s",
    },
    comment: {
      executor: "constant-vus",
      vus: 4,
      duration: "30s",
      exec: "comment",
      startTime: "12s",
    },
    postImage: {
      executor: "constant-vus",
      vus: 2,
      duration: "30s",
      exec: "postImage",
      startTime: "12s",
    }
  }
}

export default function () {}
