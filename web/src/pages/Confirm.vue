<template>
    <div>
        <div>
            <div>名前</div>
            <div>{{inquiry.name}}</div>
        </div>
        <div>
            <div>メール</div>
            <div>{{inquiry.mail}}</div>
        </div>
        <div>
            <div>問い合わせ内容</div>
            <div>{{inquiry.msg}}</div>
        </div>
        <div>
            <button @click="back">やりなおす</button>
            <button @click="submit">送信する</button>
        </div>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex';

    export default {
        name: 'Confirm',
        methods: {
            back() {
                this.$router.push('/');
            },
            submit() {
                const json = this.postInquiry();
                const router = this.$router;
                json.then(function (data) {
                    if(data.code === '200'){
                        router.push('/result');
                    }else{
                        console.log('失敗しました。');
                    }
                }).catch(function (err) {
                    console.log('エラー_' + err);
                });
            },
            async postInquiry() {
                const resp = await fetch('/api/inquiry/send', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    //TODO ここ、もっとうまくできそう
                    body: JSON.stringify({
                        'name': this.inquiry.name,
                        'mail': this.inquiry.mail,
                        'msg': this.inquiry.msg,
                    })
                });
                return await resp.json();
            },
        },
        computed: {
            ...mapGetters(['inquiry']),
        }
    }
</script>

<style scoped>

</style>