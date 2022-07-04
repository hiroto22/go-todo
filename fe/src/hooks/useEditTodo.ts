import axios from "axios";

export const useEditTodo = () => {
  const editTodo = async (token: string, id: number, todo: any) => {
    const URL = `http://127.0.0.1:8080/edittodo?id=${id}`;
    const data = { todo: todo };
    await axios
      .post(URL, JSON.stringify(data), { headers: { Authorization: token } })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  return editTodo;
};
