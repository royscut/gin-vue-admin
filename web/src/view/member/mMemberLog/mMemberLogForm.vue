
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="会员ID:" prop="memberUid">
          <el-input v-model.number="formData.memberUid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户名" />
       </el-form-item>
        <el-form-item label="操作类型:" prop="opType">
          <el-input v-model="formData.opType" :clearable="true"  placeholder="请输入操作类型" />
       </el-form-item>
        <el-form-item label="操作次数:" prop="opTimes">
          <el-input v-model.number="formData.opTimes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="付款方式:" prop="payWays">
          <el-input v-model="formData.payWays" :clearable="true"  placeholder="请输入付款方式" />
       </el-form-item>
        <el-form-item label="付款金额:" prop="payAmount">
          <el-input v-model="formData.payAmount" :clearable="true"  placeholder="请输入付款金额" />
       </el-form-item>
        <el-form-item label="续期时间:" prop="renewPeriod">
          <el-input v-model.number="formData.renewPeriod" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="对账状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="操作人ID:" prop="operatorId">
          <el-input v-model.number="formData.operatorId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="操作人姓名:" prop="operatorName">
          <el-input v-model="formData.operatorName" :clearable="true"  placeholder="请输入操作人姓名" />
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
  createMMemberLog,
  updateMMemberLog,
  findMMemberLog
} from '@/api/member/mMemberLog'

defineOptions({
    name: 'MMemberLogForm'
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
const formData = ref({
            memberUid: undefined,
            username: '',
            opType: '',
            opTimes: undefined,
            payWays: '',
            payAmount: '',
            renewPeriod: undefined,
            status: undefined,
            operatorId: undefined,
            operatorName: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               memberUid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               opType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               payWays : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               operatorId : [{
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
      const res = await findMMemberLog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createMMemberLog(formData.value)
               break
             case 'update':
               res = await updateMMemberLog(formData.value)
               break
             default:
               res = await createMMemberLog(formData.value)
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
