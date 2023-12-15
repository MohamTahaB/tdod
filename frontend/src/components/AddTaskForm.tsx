import { ChangeEvent, FormEvent } from "react"
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";


type FormProps = {
    onSubmitHandle:  (e: FormEvent<HTMLFormElement>) => Promise<void>
    onChangeHandle: (e: ChangeEvent<HTMLInputElement>) => void
}

const AddTaskForm = ({onChangeHandle, onSubmitHandle}:FormProps) => {
    return (
        <Form onSubmit={onSubmitHandle}>
            <Form.Label>Task</Form.Label>
            <Form.Control
                type="text"
                id="item"
                placeholder="Enter your task here"
                onChange={onChangeHandle}
            />
            <Button type="submit">Submit</Button>
        </Form>
    )
}

export {AddTaskForm}