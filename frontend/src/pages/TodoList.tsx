import { ChangeEvent, useEffect, useState } from "react";
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
    const [formData, setFormData] = useState<Todo>({
        completed: false,
        id: "",
        item: "",
    });

    const getData = async () => {
        try {
            const res = await api.get("/todos");
            setTodos(res.data);
        } catch (err) {
            console.log(err);
        }
    };

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value,
        })
    };

    const handleSubmit =async (e: ChangeEvent<HTMLInputElement>) => {
        try {
            await api.post("/todos", formData)
            getData()
        }
        catch (err) {
            console.log(err)
        }
    }

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
                        <tr>
                            <td>
                                <Form.Check
                                    id={task.id}
                                    checked={task.completed}
                                    onClick={() => {
                                        api.patch(`todos/${task.id}`);
                                        getData();
                                    }}
                                />
                            </td>
                            <td>{task.item}</td>
                            <td>{task.completed ? "True" : "False"}</td>
                        </tr>
                    ))}
                </tbody>
            </Table>
            <Form>
                <Form.Label>Task</Form.Label>
                <Form.Control
                    type="text"
                    id="item"
                    placeholder="Enter your task here"
                />
                <Button type="submit" onClick={() => {}}>
                    Submit
                </Button>
            </Form>
        </>
    );
}

export default TodoList;
