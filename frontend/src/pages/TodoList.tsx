import { List, Typography } from "antd";
import { useEffect, useState } from "react";
import api from "../api/todos";

function TodoList() {
    const [todos, setTodos] = useState([]);

    useEffect(() => {
        const fetchTodos = async () => {
            try {
                const res = await api.get("/todos");
                console.log("here is the raw data", res.data);
            } catch (err) {
                console.log(err);
            }
        };

        fetchTodos();
    });
    return (
        <List
            header={<div>Header</div>}
            footer={<div>Footer</div>}
            bordered
            dataSource={todos}
            renderItem={(item) => (
                <List.Item>
                    <Typography.Text mark>[ITEM]</Typography.Text> {item}
                </List.Item>
            )}
        />
    );
}

export default TodoList;
