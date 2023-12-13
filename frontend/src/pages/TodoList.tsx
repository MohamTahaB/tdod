import { useEffect, useState } from "react";
import api from "../api/todos";
import Table from "react-bootstrap/Table";
import Form from "react-bootstrap/Form";

type Todo = {
    id: string;
    item: string;
    completed: boolean;
};

function TodoList() {
    const [todos, setTodos] = useState<Todo[]>([]);

    useEffect(() => {
        const getData = async () => {
            try {
                const res = await api.get("/todos");
                setTodos(res.data);
            } catch (err) {
                console.log(err);
            }
        };
        getData();
    }, [todos]);

    return (
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
                                    setTodos([]);
                                }}
                            />
                        </td>
                        <td>{task.item}</td>
                        <td>{task.completed ? "True" : "False"}</td>
                    </tr>
                ))}
            </tbody>
        </Table>
    );
}

export default TodoList;
