import express, { Request, Response } from 'express';
import * as accountService from './services/account';
export const accountRouter = express.Router();

accountRouter.get('/', async (req: Request, res: Response) => {
  try {
    const data = await accountService.retrieve('test-user-id');

    res.status(200).send(data);
  } catch (e) {
    res.status(500).send(e.message);
  }
});

accountRouter.put('/', async (req: Request, res: Response) => {
  try {
    const data = await accountService.update('test-user-id', req.body);

    res.status(201).json(data);
  } catch (e) {
    res.status(500).send(e.message);
  }
});