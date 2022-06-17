import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { useCompleteTodo } from "../../hooks/useCompleteTodo";
import { useGetTodos } from "../../hooks/useGetTodos";
import { TodoCard } from "../molecules/todoCard";

export const Todos = () => {
  const { todos, doneTodos } = useGetTodos();
  const { CompleteTodo, ReturnTodo } = useCompleteTodo();

  return (
    <div>
      <Box sx={{ flexGrow: 1, overflow: "hidden", px: 3 }}>
        <Grid container spacing={2} wrap="nowrap">
          <Grid item xs={300} sm={1000}>
            <h1>タスク</h1>
            {todos ? (
              todos.map((res) => (
                <div>
                  <TodoCard
                    todo={res.todo}
                    text="完了"
                    onClick={() => CompleteTodo(res.id)}
                  />
                </div>
              ))
            ) : (
              <div></div>
            )}
          </Grid>
        </Grid>
        <Grid container spacing={2} wrap="nowrap">
          <Grid item xs={300} sm={1000}>
            <h1>完了</h1>
            {doneTodos ? (
              doneTodos.map((res) => (
                <div>
                  <TodoCard
                    todo={res.todo}
                    text="戻す"
                    onClick={() => ReturnTodo(res.id)}
                  />
                </div>
              ))
            ) : (
              <div></div>
            )}
          </Grid>
        </Grid>
      </Box>
    </div>
  );
};
