import axios from 'axios'
import { Todo } from '../pages/TodoList';

// Create the axios instance, to communicate with the backend. Change the baseURL according to your needs.
const api = axios.create({
    baseURL: "http://localhost:1234"
})

// GET handler:
const GetHandler: () => Promise<Todo[]> = async () => {
    try {
        const res = await api.get<Todo[]>("/todos");
        return res.data
    } catch (err) {
        console.log(err);
        throw err
    }
};

// POST handler:
const PostHandler: () => Promise<Todo[]> = async () => {
    try {
        const res = await api.get<Todo[]>("/todos");
        return res.data
    } catch (err) {
        console.log(err);
        throw err
    }
};

export {api, GetHandler}