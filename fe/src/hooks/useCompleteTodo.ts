import axios from "axios";

export const useCompleteTodo = () => {
  const CompleteTodo = (id: number) => {
    const URL = `http://127.0.0.1:8080/completetodo?id=${id}&isComplete=false`;
    axios
      .post(URL)
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  const ReturnTodo = (id: number) => {
    const URL = `http://127.0.0.1:8080/completetodo?id=${id}&isComplete=true`;
    axios
      .post(URL)
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  return { CompleteTodo, ReturnTodo };
};
