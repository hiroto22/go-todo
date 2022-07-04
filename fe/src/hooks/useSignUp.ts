import axios from "axios";
import { useNavigate } from "react-router-dom";

const URL = "http://127.0.0.1:8080/createuser";

export const useSignUp = () => {
  const navigate = useNavigate();
  const onClickSignUp = async (
    name: string | undefined,
    email: string | undefined,
    password: string | undefined
  ) => {
    const data = { name: name, email: email, password: password };
    await axios
      .post(URL, JSON.stringify(data))
      .then((res) => {
        sessionStorage.setItem("token", res.data.token);
        console.log(sessionStorage.getItem("token"));
      })
      .catch((err) => {
        console.log(err);
        alert(`${err} 
           登録できませんでした。もう一度お試し下さい。`);
      });

    navigate("/");
  };

  return onClickSignUp;
};
