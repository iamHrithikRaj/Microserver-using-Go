"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var express = require("express");
var cors = require("cors");
var app = express();
app.use(cors({
    origin: ["http://localhost:3000"]
}));
app.use(express.json());
var PORT = process.env.PORT || 8000;
app.listen(PORT, function () { return console.log("Listening on port " + PORT); });
