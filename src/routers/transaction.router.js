import Transaction from "../models/transaction.model";
import express from "express";
import {simpleRouting} from "./utils";

const router = express.Router();

simpleRouting(router, Transaction);

export default router;