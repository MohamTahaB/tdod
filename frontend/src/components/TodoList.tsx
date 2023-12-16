import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import { GetAllHandler, PostHandler, api } from "../api/todos";
import Table from "react-bootstrap/Table";
import { Todo, TodoRow } from "./TodoRow";
import { AddTaskForm } from "./AddTaskForm";

function TodoList() {
    const [todos, setTodos] = useState<Todo[]>([]);
    const [formData, setFormData] = useState<string>("");

    const getData = async () => {
        const res = await GetAllHandler();
        setTodos(res);
    };

    const deleteTaskInfo = (id : string) => {
        return async () => {try {
            const tasks = await api.delete<Todo[]>(`/todos/${id}`);
            setTodos(tasks.data)
        } catch (err) {
            console.log(err);
            throw err;
        }}
    };

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { value } = e.target;
        setFormData(value);
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        try {
            PostHandler(formData);
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
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {todos.map((task) => (
                        //TODO: define the onclickhandle
                        <TodoRow taskID={task.id} deleteTaskInfo={deleteTaskInfo} />
                    ))}
                </tbody>
            </Table>
            <AddTaskForm
                onChangeHandle={handleChange}
                onSubmitHandle={handleSubmit}
            />
        </>
    );
}

export default TodoList;
