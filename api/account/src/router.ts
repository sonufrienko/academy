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

// accountRouter.get('/:id', async (req: Request, res: Response) => {
//   const id: number = parseInt(req.params.id, 10);

//   try {
//     const item: Item = await ItemService.find(id);

//     if (item) {
//       return res.status(200).send(item);
//     }

//     res.status(404).send('item not found');
//   } catch (e) {
//     res.status(500).send(e.message);
//   }
// });

// accountRouter.post('/', async (req: Request, res: Response) => {
//   try {
//     const item: BaseItem = req.body;

//     const newItem = await ItemService.create(item);

//     res.status(201).json(newItem);
//   } catch (e) {
//     res.status(500).send(e.message);
//   }
// });

accountRouter.put('/', async (req: Request, res: Response) => {
  try {
    const data = await accountService.update('test-user-id', req.body);

    res.status(201).json(data);
  } catch (e) {
    res.status(500).send(e.message);
  }
});

// accountRouter.delete('/:id', async (req: Request, res: Response) => {
//   try {
//     const id: number = parseInt(req.params.id, 10);
//     await ItemService.remove(id);

//     res.sendStatus(204);
//   } catch (e) {
//     res.status(500).send(e.message);
//   }
// });
