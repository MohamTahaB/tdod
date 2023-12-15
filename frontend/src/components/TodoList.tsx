import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import {GetAllHandler, PostHandler, api} from "../api/todos";
import Table from "react-bootstrap/Table";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Todo, TodoRow } from "./TodoRow";

function TodoList() {
    const [todos, setTodos] = useState<Todo[]>([]);
    const [formData, setFormData] = useState<string>("");

    const getData = async () => {
        const res = await GetAllHandler()
        setTodos(res)
    };

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { value } = e.target;
        setFormData(value);
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        try {
            PostHandler(formData)
            getData();
        } catch (err) {
            console.log(err);
        }
    };

    useEffect(() => {
        getData();
    }, []);

    return (
        <>
            <Table striped bordered hover variant="dark">
                <thead>
                    <tr>
                        <th>State</th>
                        <th>Task</th>
                        <th>Done ?</th>
                    </tr>
                </thead>
                <tbody>
                    {todos.map((task) => (
                        //TODO: define the onclickhandle
                        <TodoRow task={task} onClickHandle={} />
                    ))}
                </tbody>
            </Table>
            <Form onSubmit={handleSubmit}>
                <Form.Label>Task</Form.Label>
                <Form.Control
                    type="text"
                    id="item"
                    placeholder="Enter your task here"
                    onChange={handleChange}
                />
                <Button type="submit">Submit</Button>
            </Form>
        </>
    );
}

export default TodoList;
