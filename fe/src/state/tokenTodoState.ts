import { atom } from "recoil";

export const TokenState = atom<string | null>({
  key: "bookList",
  default: "",
});
