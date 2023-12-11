import { List, Typography } from "antd";
import { useEffect, useState } from "react";
import api from "../api/todos";

type Todo = {
    id: string;
    item: string;
    completed: boolean;
};

function TodoList() {
    const [todos, setTodos] = useState<Todo[]>([]);

    useEffect(() => {
        api.get("/todos").then((res) => {
            setTodos(res.data);
        });
    }, []);
    return (
        <List
            header={<div>Header</div>}
            footer={<div>Footer</div>}
            bordered
            dataSource={todos}
            renderItem={(item) => (
                <List.Item>
                    <Typography.Text mark>
                        {item.completed ? "\uD83D\uDFE2" : "\uD83D\uDD34"}
                    </Typography.Text>{" "}
                    {item.item}
                </List.Item>
            )}
        />
    );
}

export default TodoList;
