import { create } from 'apisauce'
const createAPI = (baseURL) => {
    const api = create({
        baseURL: baseURL,
        timeout: 400,
        headers: {
        }
    });

    //groups routes
    const createGroup = async(data) => {
        const form = new FormData();
        form.append('name', data.data.name);
        return await api.post("/groups/create", form, {headers: {'content-type': 'application/json', "Authorization": data.data.token}})
    }
    const getGroups = async(data) => {
        console.log(data);
        return await api.get("/groups/list", null, {headers: {"Authorization": data.data.token}});
    }
    const joinGroup = async(data) => {
        return await api.post("/groups/groupadd/" + data.data.groupId, null, {headers:{"Authorization": data.data.token}});
    }
    const setGroup = async(data) => {
        return await api.post("groups/set/" + data.data.groupId, null, {headers:{"Authorization": data.data.token}});
    }
    const getMembers = async(data) => {
        return await api.get("/groups/users/" + data.data.groupId, null, {headers:{"Authorization": data.data.token}});
    }
    //user routes
    const login = async (data) => {
        console.log(data);
        const form = new FormData();
        form.append('email', data.data.email);
        form.append('password', data.data.password);
        return await api.post("/users/login", form, {headers: {'content-type': 'application/json', 'Access-Control-Allow-Origin': '*'}});
    }
    const register = async(data) => {
        const form = new FormData();
        form.append('email', data.data.email);
        form.append('password', data.data.password);
        form.append('username', data.data.username);
        return await api.post("/users/", form, {headers: {'content-type': 'application/json'}});
    }
    //movies routes
    const addMovie = async(data) => {
        let form = new FormData();
        form.append("name", data.data.name);
        form.append("year", data.data.year);
        return await api.post("movies/submit/" + data.data.group, form, {headers:{'content-type': 'application/json', "Authorization": data.data.token}});
    }
    const voteMovie = async(data) => {
        return await api.post("/movies/vote/" + data.data.movieId, null, {headers: {"Authorization": data.data.token}});
    }
    const getMovies = async(data) => {
        console.log("the data for get movies is:");
        console.log(data)
        return await api.get("/movies/list/" + data.data.id, null, {headers: {"Authorization": data.data.user.token}});
    }
    //nights routes
    const getNights = async(data) => {
        return await api.get("/nights/list/" + data.data.id, null, {headers:{"Authorization": data.data.user.token}});
    }
    const getHistory = async(data) => {
        return await api.get("/nights/history/" + data.data.groupId, null, {headers:{"Authorization": data.data.token}});
    }
    const addNight = async(data) => {
        let form = new FormData();
        form.append("location", data.data.location);
        form.append("date", data.data.date);
        return await api.post("/nights/new/" + data.data.group, form, {headers: {'content-type': 'application/json', "Authorization": data.data.token}});
    }
    const assignMovie = async(data) => {
        return await api.post("/nights/set/" + data.data.groupId + "/" + data.data.nightId, null, {headers: {"Authorization": data.data.token}});
    }
    return {
        getNights,
        getHistory,
        addNight,
        assignMovie,
        login,
        register,
        getGroups,
        createGroup,
        joinGroup,
        setGroup,
        getMembers,
        addMovie,
        voteMovie,
        getMovies
    };
};

export default createAPI;