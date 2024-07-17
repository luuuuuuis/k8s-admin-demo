import { createApp } from 'vue'
// 导入element plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 导入图标视图
import * as ELIcons from '@element-plus/icons-vue'
// 导入App.vue
import App from './App.vue'
// 导入路由
import router from './router'
// 导入codemirror编辑器
import { GlobalCmComponent } from "codemirror-editor-vue3"

// 创建app
const app = createApp(App)
// 将图标注册为全局组件
for (let iconName in ELIcons) {
	app.component(iconName, ELIcons[iconName])
}
// 引用element plus
app.use(ElementPlus)
// 引用codemirror编辑器
app.use(GlobalCmComponent)
// 引用路由
app.use(router)
// 挂载
app.mount('#app')
