import axios from "axios";

export const useCreateTodo = () => {
  const URL = "http://127.0.0.1:8080/createtodo";

  const createTodo = async (text: any) => {
    const data = { todo: text };
    await axios
      .post(URL, JSON.stringify(data))
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  return createTodo;
};
