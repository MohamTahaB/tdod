import axios from 'axios'
import { Todo } from '../components/TodoRow';

// Create the axios instance, to communicate with the backend. Change the baseURL according to your needs.
const api = axios.create({
    baseURL: "http://localhost:1234"
})

// GET handler: get all data.
const GetAllHandler: () => Promise<Todo[]> = async () => {
    try {
        const res = await api.get<Todo[]>("/todos");
        return res.data
    } catch (err) {
        console.log(err);
        throw err
    }
};

// POST handler: post a new task.
const PostHandler = async (item : string) => {
    try {
        await api.post("/todos", <Todo>{
            completed: false,
            id: Math.floor(Math.random() * 1000000).toString(),
            item:item,
        });
    } catch (err) {
        console.log(err);
        throw err
    }
};

export {api, GetAllHandler, PostHandler}