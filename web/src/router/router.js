import Input from "../pages/Input";
import Confirm from "../pages/Confirm";
import Result from "../pages/Result";

import  VueRouter from 'vue-router'
import Vue from 'vue';

Vue.use(VueRouter);

const routes = [
    {path: '/', component: Input},
    {path: '/confirm', component: Confirm},
    {path: '/result', component: Result}
];

export default new VueRouter({
    routes,
    mode: 'history',
});
