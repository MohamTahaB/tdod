import {
    Children,
    JSXElementConstructor,
    ReactElement,
    ReactNode,
    ReactPortal,
    useEffect,
    useState,
} from "react";
import api from "../api/todos";
import {
    Container,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
} from "@mui/material";

type Todo = {
    id: string;
    item: string;
    completed: boolean;
};

const style = {
    width: "100%",
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
    }, []);

    return (
        <Container>
            <TableContainer>
                <Table style={style}>
                    <TableHead>
                        <TableRow>
                            <TableCell align="left">Tasks</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {todos.map((todo) => (
                            <TableRow key={todo.id}>
                                <TableCell align="left">{todo.item}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </Container>
    );
}

export default TodoList;
