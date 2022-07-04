import axios from "axios";
import { useNavigate } from "react-router-dom";

const URL = "http://127.0.0.1:8080/login";

export const useLogin = () => {
  const navigate = useNavigate();
  const onClickLogin = async (
    email: string | undefined,
    password: string | undefined
  ) => {
    const data = { email: email, password: password };
    await axios
      .post(URL, JSON.stringify(data))
      .then((res) => {
        sessionStorage.setItem("token", res.data.token);
        console.log(sessionStorage.getItem("token"));
      })
      .catch((err) => {
        console.log(err);
        alert(`${err} 
         ログインできませんでした。もう一度お試し下さい。`);
      });

    navigate("/");
  };

  return onClickLogin;
};
