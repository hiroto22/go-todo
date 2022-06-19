import { TabContext, TabList, TabPanel } from "@mui/lab";
import { Grid, Tab } from "@mui/material";
import { Box } from "@mui/system";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useRecoilState } from "recoil";
import { useCompleteTodo } from "../../hooks/useCompleteTodo";
import { useGetTodos } from "../../hooks/useGetTodos";
import { editTodoState } from "../../state/editTodoState";
import { TodoCard } from "../molecules/todoCard";
import { TodoTab } from "../molecules/todoTab";

export const Todos = () => {
  const { todos, doneTodos } = useGetTodos();
  const { CompleteTodo, ReturnTodo, DeleteTodo } = useCompleteTodo();
  const [todo, setTodo] = useRecoilState(editTodoState);
  const [tabValue, setTabValue] = useState("1");
  const navigate = useNavigate();
  console.log(todos);

  return (
    <div>
      <Box sx={{ flexGrow: 1, overflow: "hidden", px: 3 }}>
        <Grid container spacing={2} wrap="nowrap">
          <Grid item xs={300} sm={1000}>
            <TabContext value={tabValue}>
              <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
                <TabList
                  onChange={(event: React.SyntheticEvent, newValue: string) =>
                    setTabValue(newValue)
                  }
                  aria-label="lab API tabs example"
                >
                  <Tab label="タスク" value="1" />
                  <Tab label="完了したタスク" value="2" />
                </TabList>
              </Box>
              <TabPanel value="1">
                {todos ? (
                  todos.map((res) => (
                    <div key={res.id}>
                      <TodoCard
                        todo={res.todo}
                        text="完了"
                        text2="削除"
                        onClick={() => CompleteTodo(res.id)}
                        onClick2={() => DeleteTodo(res.id)}
                        onClickCard={async () => {
                          await setTodo({
                            id: res.id,
                            todo: res.todo,
                            isComplete: false,
                          });
                          navigate("/edit-todo");
                        }}
                      />
                    </div>
                  ))
                ) : (
                  <div>
                    <p>タスクはありません</p>
                  </div>
                )}
              </TabPanel>
              <TabPanel value="2">
                {doneTodos ? (
                  doneTodos.map((res) => (
                    <div key={res.id}>
                      <TodoCard
                        todo={res.todo}
                        text="戻す"
                        text2="削除"
                        onClick={() => ReturnTodo(res.id)}
                        onClick2={() => DeleteTodo(res.id)}
                        onClickCard={async () => {
                          await setTodo({
                            id: res.id,
                            todo: res.todo,
                            isComplete: true,
                          });
                          navigate("/edit-todo");
                        }}
                      />
                    </div>
                  ))
                ) : (
                  <div>完了したタスクはありません</div>
                )}
              </TabPanel>
            </TabContext>
          </Grid>
        </Grid>
      </Box>
    </div>
  );
};
