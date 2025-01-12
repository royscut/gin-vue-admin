
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
      
        <el-form-item label="会员ID" prop="memberUid">
            
             <el-input v-model.number="searchInfo.memberUid" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="操作类型" prop="opType">
         <el-select v-model="searchInfo.opType" placeholder="请选择操作类型" :clearable="true" >
          <el-option label="充值卡券" value="add"></el-option>
          <el-option label="扣减卡券" value="cut"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="付款方式" prop="payWays">
         <el-select v-model="searchInfo.payWays" placeholder="请选择付款方式" :clearable="true" >
          <el-option label="现金" value="cash"></el-option>
          <el-option label="POS机" value="pos"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="对账状态" prop="status">
             <el-select v-model="searchInfo.status" placeholder="请选择对账状态" :clearable="true" >
              <el-option label="未对账" value="0"></el-option>
            <el-option label="已对账" value="1"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="操作人ID" prop="operatorId">
            
             <el-input v-model.number="searchInfo.operatorId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="操作人姓名" prop="operatorName">
         <el-input v-model="searchInfo.operatorName" placeholder="搜索条件" />
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
        <!-- <div class="gva-btn-list">
            <el-button  type="primary" icon="edit" @click="onDelete()">批量对账</el-button>
            
        </div> -->
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="会员ID" prop="memberUid" width="120" />
          <el-table-column align="left" label="用户名" prop="username" width="120" />
          <el-table-column align="left" label="操作类型" prop="opType" width="120" >
          <template #default="scope">
                {{ scope.row.opType == 'add' ? "充值卡券" : "扣减卡券" }}
            </template>
        </el-table-column>
          <el-table-column align="left" label="操作次数" prop="opTimes" width="120" />
          <el-table-column align="left" label="付款方式" prop="payWays" width="120" >
            <template #default="scope">
                {{ scope.row.payWays == 'cash' ? "现金" : scope.row.payWays == 'pos' ? "POS机" : "" }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="付款金额" prop="payAmount" width="120" />
          <el-table-column align="left" label="续期时间" prop="renewPeriod" width="120" >
            <template #default="scope">
            {{ scope.row.renewPeriod == 183 ? "半年" : scope.row.renewPeriod == 366 ? "1年" : scope.row.renewPeriod == 732 ? "2年": "" }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="对账状态" prop="status" width="120" >
            <template #default="scope">
                {{ scope.row.status == 1 ? "已对账" : "未对账" }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="操作人ID" prop="operatorId" width="120" />
          <el-table-column align="left" label="操作人姓名" prop="operatorName" width="120" />
          <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <!-- <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateMMemberLogFunc(scope.row)">编辑</el-button> -->
            <el-button  v-if="scope.row.status !== 1"  type="primary" link icon="edit" @click="checkRow(scope.row)">对账确认</el-button>
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
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="会员ID:"  prop="memberUid" >
              <el-input v-model.number="formData.memberUid" :clearable="true" placeholder="请输入会员ID" />
            </el-form-item>
            <el-form-item label="会员名:"  prop="username" >
              <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="操作类型:"  prop="opType" >
              <el-input v-model="formData.opType" :clearable="true"  placeholder="请输入操作类型" />
            </el-form-item>
            <el-form-item label="操作次数:"  prop="opTimes" >
              <el-input v-model.number="formData.opTimes" :clearable="true" placeholder="请输入操作次数" />
            </el-form-item>
            <el-form-item label="付款方式:"  prop="payWays" >
              <el-input v-model="formData.payWays" :clearable="true"  placeholder="请输入付款方式" />
            </el-form-item>
            <el-form-item label="付款金额:"  prop="payAmount" >
              <el-input v-model="formData.payAmount" :clearable="true"  placeholder="请输入付款金额" />
            </el-form-item>
            <el-form-item label="续期时间:"  prop="renewPeriod" >
              <el-input v-model.number="formData.renewPeriod" :clearable="true" placeholder="请输入续期时间" />
            </el-form-item>
            <el-form-item label="对账状态:"  prop="status" >
              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入对账状态" />
            </el-form-item>
            <el-form-item label="操作人ID:"  prop="operatorId" >
              <el-input v-model.number="formData.operatorId" :clearable="true" placeholder="请输入操作人ID" />
            </el-form-item>
            <el-form-item label="操作人姓名:"  prop="operatorName" >
              <el-input v-model="formData.operatorName" :clearable="true"  placeholder="请输入操作人姓名" />
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="会员ID">
                        {{ detailFrom.memberUid }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户名">
                        {{ detailFrom.username }}
                    </el-descriptions-item>
                    <el-descriptions-item label="操作类型">
                        {{ detailFrom.opType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="操作次数">
                        {{ detailFrom.opTimes }}
                    </el-descriptions-item>
                    <el-descriptions-item label="付款方式">
                        {{ detailFrom.payWays }}
                    </el-descriptions-item>
                    <el-descriptions-item label="付款金额">
                        {{ detailFrom.payAmount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="续期时间">
                        {{ detailFrom.renewPeriod }}
                    </el-descriptions-item>
                    <el-descriptions-item label="对账状态">
                        {{ detailFrom.status }}
                    </el-descriptions-item>
                    <el-descriptions-item label="操作人ID">
                        {{ detailFrom.operatorId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="操作人姓名">
                        {{ detailFrom.operatorName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="备注">
                        {{ detailFrom.remark }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createMMemberLog,
  deleteMMemberLog,
  deleteMMemberLogByIds,
  updateMMemberLog,
  findMMemberLog,
  getMMemberLogList
} from '@/api/member/mMemberLog'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'MMemberLog'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
              ],
               opType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               payWays : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               operatorId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
  const table = await getMMemberLogList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteMMemberLogFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要批量对账吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const data = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要对账的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          item.status = 1
          data.push(item)
        })
      
      
      const res = await updateMMemberLogList({ data })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '批量对账成功'
        })
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMMemberLogFunc = async(row) => {
    const res = await findMMemberLog({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}

// 对账确认
const checkRow = async (row) => {
    const res = await findMMemberLog({ ID: row.ID })
    if (res.code !== 0) {
      ElMessage({
                type: 'fail',
                message: '查询数据失败'
            })
    }
    if (res.data.status === 1) {
      ElMessage({
                type: 'warning',
                message: '已对账'
            })
      return
    }
    res.data.status = 1
    const res1 = await updateMMemberLog(res.data)
    if (res1.code === 0) {
        ElMessage({
                type: 'success',
                message: '对账成功'
      })
    }
    getTableData()
  }

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
                closeDialog()
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
  const res = await findMMemberLog({ ID: row.ID })
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


</script>

<style>

</style>
