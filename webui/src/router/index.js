import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from "../views/LoginView.vue";
import SuccessView from "../views/SuccessView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/success', component: SuccessView},
	]
})

export default router
