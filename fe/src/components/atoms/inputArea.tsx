type InputAreaState = {
  text: string;
  onChange: any;
};

export const InputArea = (props: InputAreaState) => {
  return (
    <div>
      <input onChange={props.onChange} placeholder={props.text}></input>
    </div>
  );
};
