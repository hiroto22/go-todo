import { atom } from "recoil";

export type EditTodoState = {
  id: number;
  todo: string;
  isComplete: boolean;
};

export const editTodoState = atom<EditTodoState>({
  key: "bookList",
  default: {
    id: 0,
    todo: "",
    isComplete: false,
  },
});
