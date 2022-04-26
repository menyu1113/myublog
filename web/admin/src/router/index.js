import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/myLogin.vue'
import Admin from '../views/myAdmin.vue'
// 页面路由
import Index from '../components/admin/myIndex.vue'
import AddArt from '../components/article/AddArt.vue'
import ArtList from '../components/article/ArtList.vue'
import CateList from '../components/category/CateList.vue'
import UserList from '../components/user/UserList.vue'
import Profile from '../components/user/myProfile.vue'
import CommentList from '../components/comment/CommentList.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    component: Login
  },
  {
    path: '/admin',
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: { title: 'Myublog' }
      },
      {
        path: 'addart',
        component: AddArt,
        meta: { title: '新增文章' }
      },
      {
        path: 'addart/:id',
        component: AddArt,
        meta: {
          title: '编辑文章'
        },
        props: true
      },
      {
        path: 'artlist',
        component: ArtList,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: 'catelist',
        component: CateList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'userlist',
        component: UserList,
        meta: {
          title: '用户列表'
        }
      },
      {
        path: 'profile',
        component: Profile,
        meta: {
          title: '个人设置'
        }
      },
      {
        path: 'commentlist',
        component: CommentList,
        meta: {
          title: '评论管理'
        }
      }
    ]

  }
]

const router = new VueRouter({
  routes
})

export default router

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()

  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    next('/login')
  } else {
    next()
  }
})
