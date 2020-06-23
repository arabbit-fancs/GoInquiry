import Vuex from 'vuex';
import Vue from 'vue';

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        inquiry:{
            name:'',
            mail:'',
            msg:'',
        }
    },
    mutations: {
        setInquiry(state, inquiry) {
            state.inquiry = inquiry
        },
        initInquiry(state) {
            state.inquiry = {
                name:'',
                mail:'',
                msg:'',
            };
        },
    },
    getters: {
        inquiry: (state) => {
            return state.inquiry
        },
    }
});