
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="会员ID:" prop="memberId">
          <el-input v-model.number="formData.memberId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总次数:" prop="totalTimes">
          <el-input v-model.number="formData.totalTimes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="剩余次数:" prop="remainTimes">
          <el-input v-model.number="formData.remainTimes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="消费次数:" prop="usedTimes">
          <el-input v-model.number="formData.usedTimes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="会员卡图片:" prop="cardPic">
          <el-input v-model="formData.cardPic" :clearable="true"  placeholder="请输入会员卡图片" />
       </el-form-item>
        <el-form-item label="过期时间:" prop="expireTime">
          <el-date-picker v-model="formData.expireTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createMMemberCard,
  updateMMemberCard,
  findMMemberCard
} from '@/api/member/mMemberCard'

defineOptions({
    name: 'MMemberCardForm'
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
            memberId: undefined,
            totalTimes: undefined,
            remainTimes: undefined,
            usedTimes: undefined,
            cardPic: '',
            expireTime: new Date(),
            remark: '',
        })
// 验证规则
const rule = reactive({
               memberId : [{
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
      const res = await findMMemberCard({ ID: route.query.id })
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
               res = await createMMemberCard(formData.value)
               break
             case 'update':
               res = await updateMMemberCard(formData.value)
               break
             default:
               res = await createMMemberCard(formData.value)
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
