import axios from "axios";
import { useEffect, useState } from "react";

type Todo = {
  id: number;
  userid: any;
  todo: string;
  createdat: Date;
  updatedat: Date;
  isDone: boolean;
};

export const useGetTodos = (token: string) => {
  const URL = "http://127.0.0.1:8080/gettodoList";
  const [todos, setTodos] = useState<Todo[]>([]);
  const [doneTodos, setDoneTodos] = useState<Todo[]>([]);

  useEffect(() => {
    console.log("postしました");
    axios
      .get(URL, { params: { isdone: 0 }, headers: { Authorization: token } })
      .then((res) => {
        setTodos(res.data);
      })
      .catch((err) => {});

    axios
      .get(URL, {
        params: { isdone: 1 },
        headers: { Authorization: token },
      })
      .then((res) => {
        setDoneTodos(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, [token]);

  return { todos, doneTodos };
};
