export interface UserData {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
}

export const retrieve = async (userId: string): Promise<UserData> => {
  return {
    id: userId,
    email: 'sergey@mail.com',
    firstName: 'Sergey',
    lastName: 'Onufrienko',
  };
};

export const update = async (userId: string, data: Partial<UserData>): Promise<UserData> => {
  const userData = await retrieve(userId);
  return userData;
};
