import * as dotenv from 'dotenv';
import express from 'express';
import cors from 'cors';
import helmet from 'helmet';
import { accountRouter } from './router';
import { errorHandler } from './middleware/error';
import { notFoundHandler } from './middleware/notFound';

dotenv.config();

if (!process.env.PORT) {
  process.exit(1);
}

const PORT: number = Number(process.env.PORT as string);

const app = express();
app.disable('x-powered-by');
app.use(helmet());
app.use(cors());
app.use(express.json());
app.use('/v1/account', accountRouter);
app.use(errorHandler);
app.use(notFoundHandler);
app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`);
});
