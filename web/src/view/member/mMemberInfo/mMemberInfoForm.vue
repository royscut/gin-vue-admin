
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="姓名:" prop="username">
          <el-input v-model="formData.username" :clearable="true"  placeholder="请输入姓名" />
       </el-form-item>
        <el-form-item label="地区:" prop="area">
           <el-select v-model="formData.area" placeholder="请选择地区" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in areaOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="手机号:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入手机号" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createMMemberInfo,
  updateMMemberInfo,
  findMMemberInfo
} from '@/api/member/mMemberInfo'

defineOptions({
    name: 'MMemberInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const areaOptions = ref([])
const formData = ref({
            username: '',
            area: '',
            mobile: '',
            address: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               username : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               area : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mobile : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMMemberInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    areaOptions.value = await getDictFunc('area')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createMMemberInfo(formData.value)
               break
             case 'update':
               res = await updateMMemberInfo(formData.value)
               break
             default:
               res = await createMMemberInfo(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
