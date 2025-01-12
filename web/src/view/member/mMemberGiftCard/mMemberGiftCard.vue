
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      <el-form-item label="礼券Key" prop="cardKey">
       <el-input v-model="searchInfo.cardKey" placeholder="=礼券Key" />
      </el-form-item>

      <el-form-item label="剩余次数" prop="remainTimes">
       <el-input v-model="searchInfo.remainTimes" placeholder=">=剩余次数" />
      </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">派发礼券</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除礼券</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="礼券Key" prop="cardKey" width="120" />
          <el-table-column align="left" label="总次数" prop="totalTimes" width="120" />
          <el-table-column align="left" label="剩余次数" prop="remainTimes" width="120" />
          <el-table-column align="left" label="核销次数" prop="usedTimes" width="120" />
         <el-table-column align="left" label="过期时间" prop="expireTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.expireTime) }}</template>
         </el-table-column>
          <el-table-column align="left" label="备注" prop="remark" width="120" />
          <el-table-column align="left" label="核销日志" prop="signOffLog" width="200" />
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看礼券</el-button>
            <el-button  v-if="scope.row.remainTimes > 0" type="primary" link icon="edit" class="table-button" @click="updateMMemberGiftCardFunc(scope.row)">核销礼券</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增礼券':'核销礼券'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <!-- <el-form-item label="礼券Key:"  prop="cardKey" >
              <el-input v-model="formData.cardKey" :clearable="true"  placeholder="请输入礼券Key" />
            </el-form-item> -->
            <el-form-item label="总次数:"  prop="totalTimes" >
              <el-input v-model.number="formData.totalTimes" :clearable="true" placeholder="请输入礼券总次数" :disabled="type !== 'create'"/>
            </el-form-item>
            <!-- <el-form-item label="剩余次数:" prop="remainTimes">
              <el-input v-model.number="formData.remainTimes" :clearable="true" placeholder="请输入剩余次数" :disabled="type !== 'create'" />
            </el-form-item> -->
            <el-form-item v-if="type !== 'create'" label="核销次数:"  prop="usedTimes" >
              <el-input v-model.number="formData.usedTimes" :clearable="true" placeholder="请输入核销次数" />
            </el-form-item>
            <!-- <el-form-item label="礼券图片:"  prop="cardPic" >
              <el-input v-model="formData.cardPic" :clearable="true"  placeholder="请输入礼券图片" />
            </el-form-item> -->
            <el-form-item label="过期时间:"  prop="expireTime" >
              <el-date-picker v-model="formData.expireTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" :disabled="type !== 'create'" />
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" :disabled="type !== 'create'"/>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogCheckFormVisible" :show-close="false" :before-close="closeCheckDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{'核销礼券'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterCheckDialog">确 定</el-button>
                  <el-button @click="closeCheckDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formCheckData" label-position="top" ref="elCheckFormRef" :rules="ruleCheck" label-width="80px">
            <!-- <el-form-item label="礼券Key:"  prop="cardKey" >
              <el-input v-model="formData.cardKey" :clearable="true"  placeholder="请输入礼券Key" />
            </el-form-item> -->
            <el-form-item label="总次数:"  prop="totalTimes" >
              <el-input v-model.number="formCheckData.totalTimes" :clearable="true" placeholder="请输入总次数" disabled />
            </el-form-item>
            <el-form-item label="剩余次数:" prop="remainTimes">
              <el-input v-model.number="formCheckData.remainTimes" :clearable="true" placeholder="请输入剩余次数" disabled />
            </el-form-item>
            <el-form-item label="本次核销次数:"  prop="usedTimes" >
              <el-input v-model.number="formCheckData.usedTimes" :clearable="true" placeholder="请输入本次核销次数" />
            </el-form-item>
            <!-- <el-form-item label="卡券图片:"  prop="cardPic" >
              <el-input v-model="formData.cardPic" :clearable="true"  placeholder="请输入卡券图片" />
            </el-form-item> -->
            <el-form-item label="过期时间:"  prop="expireTime" >
              <el-date-picker v-model="formCheckData.expireTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" disabled />
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formCheckData.remark" :clearable="true"  placeholder="请输入备注" disabled/>
            </el-form-item>
            <el-form-item label="核销日志:"  prop="signOffLog" >
              <el-input v-model="formCheckData.signOffLog" :clearable="true"  placeholder="" disabled/>
            </el-form-item>
          </el-form>
    </el-drawer>


    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="礼券Key">
                        {{ detailFrom.cardKey }}
                    </el-descriptions-item>
                    <el-descriptions-item label="总次数">
                        {{ detailFrom.totalTimes }}
                    </el-descriptions-item>
                    <el-descriptions-item label="剩余次数">
                        {{ detailFrom.remainTimes }}
                    </el-descriptions-item>
                    <el-descriptions-item label="核销次数">
                        {{ detailFrom.usedTimes }}
                    </el-descriptions-item>
                    
                    <el-descriptions-item label="过期时间">
                        {{ detailFrom.expireTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="备注">
                        {{ detailFrom.remark }}
                    </el-descriptions-item>

                    <el-descriptions-item label="礼券图片">
                      <img :src="detailFrom.cardPic" alt="礼券图片" style="width: 100px; height: auto; cursor: pointer;" @click="viewImage(detailFrom.cardPic)" />
                <div>
                  <el-button type="text" @click="viewImage(detailFrom.cardPic)">查看原图</el-button>
                </div>
                    </el-descriptions-item>
            </el-descriptions>
          </el-drawer>
  </div>
</template>

<script setup>
import {
  createMMemberGiftCard,
  deleteMMemberGiftCard,
  deleteMMemberGiftCardByIds,
  updateMMemberGiftCard,
  findMMemberGiftCard,
  getMMemberGiftCardList
} from '@/api/member/mMemberGiftCard'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'MMemberGiftCard'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            cardKey: '',
            totalTimes: undefined,
            remainTimes: undefined,
            usedTimes: undefined,
            cardPic: '',
            expireTime: new Date(),
            remark: '',
        })

const formCheckData = ref({
            cardKey: '',
            totalTimes: undefined,
            remainTimes: undefined,
            usedTimes: undefined,
            cardPic: '',
            expireTime: new Date(),
            remark: '',
            signOffLog: '',
        })



// 验证规则
const rule = reactive({
})

// 验证规则
const ruleCheck = reactive({
  usedTimes: [
    { 
      required: true, 
      message: '请输入核销次数', 
      trigger: ['input', 'blur'],
      type: 'number'
    },
    { 
      type: 'number', 
      min: 1, 
      message: '核销次数不能小于0', 
      trigger: ['input', 'blur']
    }
  ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elCheckFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMMemberGiftCardList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMMemberGiftCardFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('只有未核销过的礼券才可删除，确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteMMemberGiftCardByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMMemberGiftCardFunc = async(row) => {
    const res = await findMMemberGiftCard({ ID: row.ID })
    type.value = 'check'
    if (res.code === 0) {
      res.data.usedTimes = undefined
        formCheckData.value = res.data
        dialogCheckFormVisible.value = true
    }
}


// 删除行
const deleteMMemberGiftCardFunc = async (row) => {
    const res = await deleteMMemberGiftCard({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)
const dialogCheckFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    formData.value.totalTimes = 1
    dialogFormVisible.value = true

}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        cardKey: '',
        totalTimes: undefined,
        remainTimes: undefined,
        usedTimes: undefined,
        cardPic: '',
        expireTime: new Date(),
        remark: '',
        }
}

// 关闭弹窗
const closeCheckDialog = () => {
    dialogCheckFormVisible.value = false
    formCheckData.value = {
        cardKey: '',
        totalTimes: undefined,
        remainTimes: undefined,
        usedTimes: undefined,
        cardPic: '',
        expireTime: new Date(),
        remark: '',
        signOffLog: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createMMemberGiftCard(formData.value)
                  break
                case 'update':
                  res = await updateMMemberGiftCard(formData.value)
                  break
                default:
                  res = await createMMemberGiftCard(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '派发礼券成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

// 弹窗确定
const enterCheckDialog = async () => {
     btnLoading.value = true
     elCheckFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'check':
                  res = await updateMMemberGiftCard(formCheckData.value)
                  break
                default:
                  res = await updateMMemberGiftCard(formCheckData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '核销礼券成功'
                })
                closeCheckDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findMMemberGiftCard({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

const viewImage = (imageUrl) =>{
      window.open(imageUrl, '_blank');
}


</script>

<style>

</style>
