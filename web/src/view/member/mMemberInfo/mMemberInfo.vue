
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <!-- <el-form-item label="创建日期" prop="createdAt">
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
      </el-form-item> --> 
      <el-form-item label="会员ID" prop="memberId">
            
            <el-input v-model.number="searchInfo.memberId" placeholder="搜索条件" />
       </el-form-item>

        <el-form-item label="姓名" prop="username">
         <el-input v-model="searchInfo.username" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="手机号" prop="mobile">
         <el-input v-model="searchInfo.mobile" placeholder="搜索条件" />
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
            <el-button  type="primary" icon="plus" @click="openDialog()">新增会员</el-button>
            <!-- <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button> -->
            
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
        
        <el-table-column align="left" label="日期" prop="createdAt"width="150">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="会员ID" prop="ID" width="100" />
          <el-table-column align="left" label="姓名" prop="username" width="100" />
        <el-table-column align="left" label="地区" prop="area" width="80">
            <template #default="scope">
                {{ filterDict(scope.row.area,areaOptions) }}
            </template>
        </el-table-column>
          <el-table-column align="left" label="手机号" prop="mobile" width="120" />
          <el-table-column align="left" label="住址" prop="address" width="120" />
      
          <el-table-column align="left" label="总次数" prop="totalTimes" width="80" />
          <el-table-column align="left" label="剩余次数" prop="remainTimes" width="80" />
          <el-table-column align="left" label="消费次数" prop="usedTimes" width="80" />
         <el-table-column align="left" label="过期时间" prop="expireTime" width="150">
            <template #default="scope">{{ formatDate(scope.row.expireTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="备注" prop="remark" width="120" />

         <el-table-column align="left" label="查看"  width="210">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>会员详情</el-button>
            <el-button  type="primary" link class="table-button" @click="goToH5Detail(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看H5</el-button>
            </template>
          </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateMMemberInfoFunc(scope.row)">编辑</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateMMemberCardFunc(scope.row,'add')">充卡</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateMMemberCardFunc(scope.row,'cut')">扣卡</el-button>
            <!-- <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button> -->

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
                <span class="text-lg">{{type==='create'?'新增会员信息':'编辑会员信息'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item v-if="type !== 'create'" label="会员ID:"  prop="ID" >
              <el-input v-model.number="formData.ID"  placeholder="请输入会员ID" disabled/>
            </el-form-item>
            <el-form-item label="用户名:"  prop="username" >
            <el-input v-model="formData.username" :clearable="true"  placeholder="请输入姓名" />
            </el-form-item>
            <el-form-item label="地区:"  prop="area" >
              <el-select v-model="formData.area" placeholder="请选择地区" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in areaOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="手机号:"  prop="mobile" >
              <el-input v-model="formData.mobile" :clearable="true"  placeholder="请输入手机号" />
            </el-form-item>
            <el-form-item label="地址:"  prop="address" >
              <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogCardFormVisible" :show-close="false" :before-close="closeCardDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='add'?'充值卡券':'扣减卡券'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterCardDialog">确 定</el-button>
                  <el-button @click="closeCardDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formcardData" label-position="top" ref="elCardFormRef" :rules="cardRule" label-width="80px">
            <el-form-item label="会员ID:"  prop="memberUid" >
              <el-input v-model.number="formcardData.memberUid" :clearable="true" placeholder="请输入会员ID" disabled/>
            </el-form-item>
            <el-form-item label="用户名:"  prop="username" >
              <el-input v-model="formcardData.username" :clearable="true"  placeholder="请输入用户名" disabled/>
            </el-form-item>
            <el-form-item label="操作类型:"  prop="opType">
              <span>{{type==='add'?'充值卡券':'扣减卡券'}}</span>
              <el-input v-if="false" v-model="formcardData.opType" :clearable="true"  placeholder="请输入操作类型" />
            </el-form-item>
            <el-form-item v-if="type === 'add'"label="充值次数:"  prop="opTimes" >
              <el-input v-model.number="formcardData.opTimes" :clearable="true" placeholder="请输入充值次数" />
            </el-form-item>

            <el-form-item v-if="type === 'cut'" label="扣减次数:"  prop="opTimes" >
              <el-input v-model.number="formcardData.opTimes" :clearable="true" placeholder="请输入扣减次数" />
            </el-form-item>

            <el-form-item v-if="type === 'add'" label="付款方式:" prop="payWays">
              <el-radio-group v-model="formcardData.payWays">
              <el-radio label="cash">现金</el-radio>
              <el-radio label="pos">POS机</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item v-if="type === 'add'" label="付款金额:"  prop="payAmount" >
              <el-input v-model="formcardData.payAmount" :clearable="true"  placeholder="请输入付款金额" />
            </el-form-item>
            <el-form-item v-if="type === 'add'" label="续期时间:"  prop="renewPeriod" >
              <el-radio-group v-model="formcardData.renewPeriod">
              <el-radio label="183">半年(183天)</el-radio>
              <el-radio label="366">1年(366天)</el-radio>
              <el-radio label="732">2年(732天)</el-radio>
              </el-radio-group>
            </el-form-item>

            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formcardData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看会员卡信息">
            <el-descriptions :column="1" border>
                  <el-descriptions-item label="会员ID">
                        {{ detailFrom.ID }}
                    </el-descriptions-item>
                    <el-descriptions-item label="姓名">
                        {{ detailFrom.username }}
                    </el-descriptions-item>
                    <el-descriptions-item label="地区">
                        
                        {{ filterDict(detailFrom.area,areaOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="手机号">
                        {{ detailFrom.mobile }}
                    </el-descriptions-item>
                    <el-descriptions-item label="地址">
                        {{ detailFrom.address }}
                    </el-descriptions-item>
                    <el-descriptions-item label="备注">
                        {{ detailFrom.remark }}
                    </el-descriptions-item>
                    <el-descriptions-item label="礼券图片">
                      <img :src="detailFrom.cardPic" alt="卡券图片" style="width: 100px; height: auto; cursor: pointer;" @click="viewImage(detailFrom.cardPic)" />
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
  createMMemberInfo,
  deleteMMemberInfo,
  deleteMMemberInfoByIds,
  updateMMemberInfo,
  findMMemberInfo,
  getMMemberInfoList
} from '@/api/member/mMemberInfo'

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
    name: 'MMemberInfo'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const areaOptions = ref([])
const formData = ref({
            username: '',
            area: '',
            mobile: '',
            address: '',
            remark: '',
        })

const formcardData = ref({
            memberUid: undefined,
            username: '',
            opType: '',
            opTimes: undefined,
            payWays: '',
            payAmount: '',
            renewPeriod: undefined,
            remark: '',
        })



// 验证规则
const rule = reactive({
              username : [{
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
               area : [{
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
               mobile : [{
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
})

const carRule = reactive({
              username : [{
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

              memberUid : [{
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

              opTimes : [{
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
const elCardFormRef = ref()

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
  const table = await getMMemberInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    areaOptions.value = await getDictFunc('area')
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
            deleteMMemberInfoFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
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
      const res = await deleteMMemberInfoByIds({ IDs })
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
const updateMMemberInfoFunc = async(row) => {
    const res = await findMMemberInfo({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}

// 充卡和扣卡,1为充卡，2为扣卡  
const updateMMemberCardFunc = async(row,opType) => {
    const res = await findMMemberInfo({ ID: row.ID })
    if (opType === 'add') {
        type.value = 'add'
    } else {
        type.value = 'cut'
    }

    if (res.code === 0) {
      formcardData.value = {
            memberUid: res.data.ID,
            username: res.data.username,
            opType: opType,
            opTimes: undefined,
            payWays: 'cash',
            payAmount: '',
            renewPeriod: "183",
            remark: '',
        }
        dialogCardFormVisible.value = true
    }
}


// 删除行
const deleteMMemberInfoFunc = async (row) => {
    const res = await deleteMMemberInfo({ ID: row.ID })
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
// 弹窗控制标记
const dialogCardFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    formData.value = {
        username: '',
        area: 'hk',
        mobile: '',
        address: '',
        remark: '',
        }
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        username: '',
        area: '',
        mobile: '',
        address: '',
        remark: '',
        }
}

// 关闭弹窗
const closeCardDialog = () => {
  dialogCardFormVisible.value = false
    formcardData.value = {
        username: '',
        area: '',
        mobile: '',
        address: '',
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
                closeDialog()
                getTableData()
              }
      })
}


// 弹窗确定
const enterCardDialog = async () => {
     btnLoading.value = true
     elCardFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              let msg = ''
              switch (type.value) {
                case 'add':
                  msg = '充值卡券成功'
                  break
                case 'cut':
                  msg = '扣减卡券成功'
                  formcardData.value.renewPeriod = undefined
                  formcardData.value.payAmount = ''
                  formcardData.value.payWays = ''
                  break
                default:
                   msg = '充值卡券成功'
                  break
              }
              
              res = await createMMemberLog(formcardData.value)

              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: msg
                })
                closeCardDialog()
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
  const res = await findMMemberInfo({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}

const goToH5Detail = (row) => {
  // 方式一：新窗口打开
  const url = `#/h5/member/detail?id=${row.ID}`
  window.open(url, '_blank')
  
  // 或者方式二：使用路由跳转
  // router.push({
  //   path: '/h5/member/detail',
  //   query: { id: row.ID }
  // })
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
