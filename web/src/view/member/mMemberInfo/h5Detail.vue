<template>
    <div class="member-detail">
      <!-- 顶部导航栏 -->
      <div class="nav-header">
        <div class="back" @click="goBack">
          <el-icon><ArrowLeft /></el-icon>
        </div>
        <div class="title">会员卡详情</div>
        <div class="placeholder"></div>
      </div>
  
    
      <!-- 详情内容 -->
      <div class="detail-content">
           <!-- 消费信息卡片 -->
           <div class="times-card">
          <div class="times-item">
            <div class="num">{{ detailInfo.totalTimes }}</div>
            <div class="text">总次数</div>
          </div>
          <div class="times-item">
            <div class="num">{{ detailInfo.remainTimes }}</div>
            <div class="text">剩余次数</div>
          </div>
          <div class="times-item">
            <div class="num">{{ detailInfo.usedTimes }}</div>
            <div class="text">已用次数</div>
          </div>
        </div>

        <div class="info-item">
          <div class="label">会员ID
          </div>
          <div class="value">{{ detailInfo.ID }}</div>
        </div>
        
        <div class="info-item">
          <div class="label">用户名</div>
          <div class="value">{{ detailInfo.username }}</div>
        </div>
  
        <div class="info-item">
          <div class="label">地区</div>
          <div class="value">{{ filterDict(detailInfo.area, areaOptions) }}</div>
        </div>
  
        <div class="info-item">
          <div class="label">手机号</div>
          <div class="value">{{ detailInfo.mobile }}</div>
        </div>
  
        <div class="info-item">
          <div class="label">地址</div>
          <div class="value">{{ detailInfo.address }}</div>
        </div>
  
       
  
        <div class="info-item remark">
          <div class="label">备注</div>
          <div class="value">{{ detailInfo.remark || '-' }}</div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { ArrowLeft } from '@element-plus/icons-vue'
  import { findMMemberInfo } from '@/api/member/mMemberInfo'
  import { filterDict } from '@/utils/format'
  
  const route = useRoute()
  const router = useRouter()
  const detailInfo = ref({})
  const areaOptions = ref([])
  
  // 获取详情数据
  const getDetail = async () => {
    const res = await findMMemberInfo({ ID: route.query.id })
    if (res.code === 0) {
      detailInfo.value = res.data
    }
  }
  
  const goBack = () => {
    router.back()
  }
  
  onMounted(() => {
    getDetail()
  })
  </script>
  
  <style lang="scss" scoped>
  .member-detail {
    min-height: 100vh;
    background: #f5f5f5;
  
    .nav-header {
      position: sticky;
      top: 0;
      z-index: 100;
      display: flex;
      align-items: center;
      justify-content: space-between;
      height: 44px;
      padding: 0 16px;
      background: #fff;
      box-shadow: 0 1px 2px rgba(0,0,0,0.1);
  
      .back {
        padding: 8px;
        margin-left: -8px;
      }
  
      .title {
        font-size: 16px;
        font-weight: 500;
      }
  
      .placeholder {
        width: 24px;
      }
    }
  
    .detail-content {
      padding: 16px;
  
      .info-item {
        margin-bottom: 16px;
        padding: 16px;
        background: #fff;
        border-radius: 8px;
  
        .label {
          color: #666;
          font-size: 14px;
          margin-bottom: 8px;
        }
  
        .value {
          color: #333;
          font-size: 15px;
        }
  
        &.remark {
          .value {
            white-space: pre-wrap;
            line-height: 1.5;
          }
        }
      }
  
      .times-card {
        display: flex;
        justify-content: space-between;
        padding: 20px 16px;
        background: #fff;
        border-radius: 8px;
        margin-bottom: 16px;
  
        .times-item {
          text-align: center;
          flex: 1;
  
          .num {
            font-size: 24px;
            color: var(--el-color-primary);
            font-weight: 500;
            margin-bottom: 4px;
          }
  
          .text {
            font-size: 13px;
            color: #666;
          }
        }
      }
    }
  }
  </style>