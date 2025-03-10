// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	scyllav1alpha1 "github.com/scylladb/scylla-operator/pkg/client/scylla/clientset/versioned/typed/scylla/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeScyllaDBClusters implements ScyllaDBClusterInterface
type fakeScyllaDBClusters struct {
	*gentype.FakeClientWithList[*v1alpha1.ScyllaDBCluster, *v1alpha1.ScyllaDBClusterList]
	Fake *FakeScyllaV1alpha1
}

func newFakeScyllaDBClusters(fake *FakeScyllaV1alpha1, namespace string) scyllav1alpha1.ScyllaDBClusterInterface {
	return &fakeScyllaDBClusters{
		gentype.NewFakeClientWithList[*v1alpha1.ScyllaDBCluster, *v1alpha1.ScyllaDBClusterList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("scylladbclusters"),
			v1alpha1.SchemeGroupVersion.WithKind("ScyllaDBCluster"),
			func() *v1alpha1.ScyllaDBCluster { return &v1alpha1.ScyllaDBCluster{} },
			func() *v1alpha1.ScyllaDBClusterList { return &v1alpha1.ScyllaDBClusterList{} },
			func(dst, src *v1alpha1.ScyllaDBClusterList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.ScyllaDBClusterList) []*v1alpha1.ScyllaDBCluster {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.ScyllaDBClusterList, items []*v1alpha1.ScyllaDBCluster) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
