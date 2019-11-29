# encoding: utf-8

class test:
    def test_main(self, args):
        # print('main')
        print(args)
        env_ip = args.get('env_ip')
        version = args.get('version')
        infra_ansible_branch = args.get('infra_ansible')
        tools_branch = args.get('tools')
        print('python main.py %s --version=%s --infra_ansible=%s tools=%s' % (
            env_ip, version, infra_ansible_branch, tools_branch))
        return ('python main.py %s --version=%s --infra_ansible=%s tools=%s' % (
            env_ip, version, infra_ansible_branch, tools_branch))


t1 = test()
