import { TabsContext } from "@mui/base";
import { PropaneSharp } from "@mui/icons-material";
import { TabContext, TabList, TabPanel } from "@mui/lab";
import { Tab } from "@mui/material";
import { Box } from "@mui/system";

type TabState = {
  value: string;
  onChange: any;
  content: any;
};

export const TodoTab = (props: TabState) => {
  return (
    <TabContext value={props.value}>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <TabList onChange={props.onChange} aria-label="lab API tabs example">
          <Tab label="タスク" value="1" />
          <Tab label="完了したタスク" value="2" />
        </TabList>
      </Box>
      <TabPanel value="1">{props.content}</TabPanel>
      <TabPanel value="2">Item Two</TabPanel>
    </TabContext>
  );
};
