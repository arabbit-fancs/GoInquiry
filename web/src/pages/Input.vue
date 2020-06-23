<template>
    <div>
        <div>
            <div>
                お名前
            </div>
            <input type="text" v-model="inputInquiry.name"/>
        </div>
        <div>
            <div>
                メールアドレス
            </div>
            <input type="text" v-model="inputInquiry.mail"/>
        </div>
        <div>
            <div>
                問い合わせ内容
            </div>
            <input type="text" v-model="inputInquiry.msg"/>
        </div>
        <div>
            <button @click="onClick">確認画面へ</button>
        </div>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex';

    export default {
        name: 'Input',
        data() {
            return {
                inputInquiry: {
                    name: '',
                    mail: '',
                    msg: '',
                },
            }
        },
        created(){
          this.inputInquiry =  this.inquiry;
        },
        methods: {
            onClick() {
                if (this.inputInquiry.name !== '' && this.inputInquiry.mail !== '' && this.inputInquiry.msg !== '') {
                    this.$store.commit('setInquiry',this.inputInquiry);
                    this.$router.push('/confirm');
                }
            }
        },
        computed: {
            ...mapGetters(['inquiry']),
        }
    }
</script>

<style scoped>

</style>