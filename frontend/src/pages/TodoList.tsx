import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import api from "../api/todos";
import Table from "react-bootstrap/Table";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

type Todo = {
    id: string;
    item: string;
    completed: boolean;
};

function TodoList() {
    const [todos, setTodos] = useState<Todo[]>([]);
    const [formData, setFormData] = useState<string>("");

    const getData = async () => {
        try {
            const res = await api.get("/todos");
            setTodos(res.data);
        } catch (err) {
            console.log(err);
        }
    };

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { value } = e.target;
        setFormData(value);
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        try {
            await api.post("/todos", {
                completed: false,
                id: Math.floor(Math.random() * 1000000).toString(),
                item: formData,
            });
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
                        <tr id={task.id}>
                            <td>
                                <Form.Check
                                    id={task.id}
                                    checked={task.completed}
                                    onClick={async () => {
                                        await api.patch(`todos/${task.id}`);
                                        await getData();
                                    }}
                                />
                            </td>
                            <td>{task.item}</td>
                            <td>{task.completed ? "True" : "False"}</td>
                        </tr>
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
