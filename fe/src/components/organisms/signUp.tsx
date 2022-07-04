import { Card, CardContent, Grid } from "@mui/material";
import { useState } from "react";
import { useSignUp } from "../../hooks/useSignUp";
import { BaseButton } from "../atoms/baseButton";
import { InputArea } from "../atoms/inputArea";

export const SignUp = () => {
  const [name, setName] = useState<string>();
  const [email, setEmail] = useState<string>();
  const [passWord, setPassWord] = useState<string>();
  const onClickSignUp = useSignUp();

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
              <h1>ユーザー登録</h1>
            </Grid>
            <Grid item justifyContent="center">
              <p>
                登録済みの方は<a href="http://localhost:3000/login">こちら</a>へ
              </p>
            </Grid>
            <Grid item>
              <InputArea
                text="名前"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                  setName(e.target.value)
                }
              />
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
                onClick={() => onClickSignUp(name, email, passWord)}
                text={"新規登録"}
              />
            </Grid>
          </Grid>
        </CardContent>
      </Card>
    </Grid>
  );
};
