import styled from "styled-components";

type TextAreaState = {
  defaultValue: string;
  onChange: any;
};

export const TextArea = (props: TextAreaState) => {
  return (
    <STextArea
      onChange={props.onChange}
      defaultValue={props.defaultValue}
    ></STextArea>
  );
};

const STextArea = styled.textarea`
  width: 99%;
  height: 60px;
  display: flex;
  margin-left: auto;
  margin-right: auto;
`;
