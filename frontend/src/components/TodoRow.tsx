import React, { useEffect, useState } from "react";
import Form from "react-bootstrap/esm/Form";
import { api } from "../api/todos";
import Button from "react-bootstrap/esm/Button";

type Todo = {
    id: string;
    item: string;
    completed: boolean;
};

type TodoRowProps = {
    taskID: string;
};

const TodoRow = ({ taskID }: TodoRowProps) => {
    const [task, setTask] = useState<Todo>({
        completed: false,
        id: "",
        item: "",
    });

    const fetchTaskInfo = async () => {
        try {
            const taskInfo = await api.get(`/todos/${taskID}`);
            setTask(taskInfo.data);
        } catch (err) {
            console.log(err);
            throw err;
        }
    };

    const patchTaskInfo = async () => {
        try {
            const newTask: Todo = {
                ...task,
                completed: done.valueOf(),
            };
            await api.patch(`/todos/${taskID}`);
        } catch (err) {
            console.log(err);
            throw err;
        }
    };

    useEffect(() => {
        fetchTaskInfo();
    }, []);

    const [done, setDone] = useState<Boolean>(task.completed);

    useEffect(() => {
        patchTaskInfo();
        fetchTaskInfo();
    }, [done]);

    return (
        <tr id={task.id}>
            <td>
                <Form.Check
                    id={task.id}
                    checked={task.completed}
                    onClick={() => {
                        setDone(!done);
                    }}
                />
            </td>
            <td>{task.item}</td>
            <td>{task.completed ? "True" : "False"}</td>
            <td>
                <Button variant="danger">
                    Delete
                </Button>
            </td>
        </tr>
    );
};

export { TodoRow };
export type { Todo };
