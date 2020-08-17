import { create } from 'apisauce'
const createAPI = (baseURL) => {
    const api = create({
        baseURL: baseURL,
        timeout: 400,
        headers: {
        }
    });

    //groups routes
    const getMovie = async(data) => {
        return await api.get("/", {...data, apikey: "fdafcd67"});
    }

    return {
        getMovie
    };
};

export default createAPI;