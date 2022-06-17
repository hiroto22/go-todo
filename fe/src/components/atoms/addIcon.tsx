import AddCircleIcon from "@mui/icons-material/AddCircle";
import styled from "styled-components";

export const AddIcon = () => {
  return (
    <SDiv>
      <AddCircleIcon sx={{ width: "10%", height: "10%" }} />
    </SDiv>
  );
};

const SDiv = styled.div`
  color: #0277bd;
  display: flex;
  justify-content: right;
  margin-right: 7%;
`;
