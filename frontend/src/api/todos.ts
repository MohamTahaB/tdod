import axios from 'axios'
import { Todo } from '../pages/TodoList';

// Create the axios instance, to communicate with the backend. Change the baseURL according to your needs.
const api = axios.create({
    baseURL: "http://localhost:1234"
})

// GET handler:
// - Gets the data, and sets a hook state
const getData: Promise<Todo[]> = async () => {
    try {
        const res = await api.get("/todos");
        return res.data typeof Todo[]
    } catch (err) {
        console.log(err);
    }
};

export {api}