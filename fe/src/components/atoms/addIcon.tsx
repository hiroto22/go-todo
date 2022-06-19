import AddCircleIcon from "@mui/icons-material/AddCircle";
import styled from "styled-components";

type AddIconBUtton = {
  onClick: any;
};

export const AddIcon = (props: AddIconBUtton) => {
  return (
    <SDiv>
      <AddCircleIcon
        sx={{ width: "10%", height: "10%" }}
        onClick={props.onClick}
      />
    </SDiv>
  );
};

const SDiv = styled.div`
  color: #0277bd;
  display: flex;
  justify-content: right;
  margin-right: 7%;
`;
