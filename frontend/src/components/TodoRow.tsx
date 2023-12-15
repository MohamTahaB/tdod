import React from "react";
import Form from "react-bootstrap/esm/Form";
import { api } from "../api/todos";

type Todo = {
    id: string;
    item: string;
    completed: boolean;
};

type TodoRowProps = {
    task : Todo
    onClickHandle : () => Promise<void>
}

const TodoRow = ({task, onClickHandle} : TodoRowProps) => {
    return (
        <tr id={task.id}>
            <td>
                <Form.Check
                    id={task.id}
                    checked={task.completed}
                    onClick={onClickHandle}
                />
            </td>
            <td>{task.item}</td>
            <td>{task.completed ? "True" : "False"}</td>
        </tr>
    )
}

export {TodoRow}
export type {Todo}