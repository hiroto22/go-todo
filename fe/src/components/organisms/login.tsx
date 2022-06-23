import { Card, CardContent, Grid } from "@mui/material";
import { useState } from "react";
import { useLogin } from "../../hooks/useLogin";
import { BaseButton } from "../atoms/baseButton";
import { InputArea } from "../atoms/inputArea";

export const Login = () => {
  const [email, setEmail] = useState<string>();
  const [passWord, setPassWord] = useState<string>();
  const onClickLogin = useLogin();

  return (
    <Grid container justifyContent="center" padding={10}>
      <Card sx={{ width: 800, minHeight: 400 }}>
        <CardContent>
          <Grid
            container
            spacing={1}
            direction="column"
            justifyContent="center"
            alignItems="center"
          >
            <Grid item justifyContent="center">
              <h1>ログイン</h1>
            </Grid>
            <Grid item justifyContent="center">
              <p>
                新規登録の方は<a href="http://localhost:3000/signup">こちら</a>
                へ
              </p>
            </Grid>
            <Grid item>
              <InputArea
                text="メールアドレス"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                  setEmail(e.target.value)
                }
              />
            </Grid>
            <Grid item>
              <InputArea
                text="パスワード"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                  setPassWord(e.target.value)
                }
              />
            </Grid>
            <Grid item>
              <BaseButton
                onClick={() => onClickLogin(email, passWord)}
                text={"ログイン"}
              />
            </Grid>
          </Grid>
        </CardContent>
      </Card>
    </Grid>
  );
};
